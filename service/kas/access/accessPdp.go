package access
import (
    "context"
    "errors"
    "slices"
    "strings"
    "time"
    "github.com/opentdf/platform/protocol/go/authorization"
    "github.com/opentdf/platform/protocol/go/policy"
    otdf "github.com/opentdf/platform/sdk"
    "github.com/opentdf/platform/service/logger"
)
const (
    ErrPolicyDissemInvalid     = Error("policy dissem invalid")
    ErrDecisionUnexpected      = Error("authorization decision unexpected")
    ErrDecisionCountUnexpected = Error("authorization decision count unexpected")
)
func canAccess(ctx context.Context, token *authorization.Token, policy Policy, sdk *otdf.SDK, logger logger.Logger) (bool, error) {
    if len(policy.Body.Dissem) > 0 {
        // TODO: Move dissems check to the getdecisions endpoint
        logger.Error("Dissems check is not enabled in v2 platform kas")
    }
    if policy.Body.DataAttributes != nil {
        attrAccess, err := checkAttributes(ctx, policy.Body.DataAttributes, token, sdk, logger)
        if err != nil {
            return false, err
        }
        return attrAccess, nil
    }
    // if no dissem and no attributes then allow
    return true, nil
}
func checkAttributes(ctx context.Context, dataAttrs []Attribute, ent *authorization.Token, sdk *otdf.SDK, logger logger.Logger) (bool, error) {
    ras := []*authorization.ResourceAttribute{{
        AttributeValueFqns: make([]string, 0),
    }}
    for _, attr := range dataAttrs {
        if attr.Name != "temporal" {
            ras[0].AttributeValueFqns = append(ras[0].GetAttributeValueFqns(), attr.URI)
        }
    }
    for _, attr := range dataAttrs {
        attrSplit := strings.Split(attr.ProviderURI, "/")
        // for example: "https://demo.com/attr/notbefore/value/2013-Feb-03"
        if attr.Name == "temporal" {
            tNow := time.Now().UTC()
            const shortForm = "2006-Jan-02"
            tTdf, _ := time.Parse(shortForm, attrSplit[len(attrSplit)-1])
            if slices.Contains(attrSplit, "notbefore") && tTdf.After(tNow) {
                return false, nil
            } else if slices.Contains(attrSplit, "notafter") && tNow.After(tTdf) {
                return false, nil
            }
        }
    }
    in := authorization.GetDecisionsByTokenRequest{
        DecisionRequests: []*authorization.TokenDecisionRequest{
            {
                Actions: []*policy.Action{
                    {Value: &policy.Action_Standard{Standard: policy.Action_STANDARD_ACTION_DECRYPT}},
                },
                Tokens:             []*authorization.Token{ent},
                ResourceAttributes: ras,
            },
        },
    }
    dr, err := sdk.Authorization.GetDecisionsByToken(ctx, &in)
    if err != nil {
        logger.ErrorContext(ctx, "Error received from GetDecisionsByToken", "err", err)
        return false, errors.Join(ErrDecisionUnexpected, err)
    }
    if len(dr.GetDecisionResponses()) != 1 {
        logger.ErrorContext(ctx, ErrDecisionCountUnexpected.Error(), "count", len(dr.GetDecisionResponses()))
        return false, ErrDecisionCountUnexpected
    }
    if dr.GetDecisionResponses()[0].GetDecision() == authorization.DecisionResponse_DECISION_PERMIT {
        return true, nil
    }
    return false, nil
}