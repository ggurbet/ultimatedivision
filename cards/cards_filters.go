package cards

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/sqlsearchoperators"
)

const (
	// numberPositionOfURLParameter is a number that shows the position of the url parameter.
	numberPositionOfURLParameter = 0
)

// ErrInvalidFilter indicated that filter does not valid.
var ErrInvalidFilter = errs.Class("invalid filter")

// Filters entity for using filter cards.
type Filters struct {
	Name           Filter
	Value          string
	SearchOperator sqlsearchoperators.SearchOperator
}

// Filter defines the list of possible filters.
type Filter string

const (
	// FilterTactics indicates filtering by card tactics.
	FilterTactics Filter = "tactics"
	// FilterPositioning indicates filtering by card positioning.
	FilterPositioning Filter = "positioning"
	// FilterComposure indicates filtering by card composure.
	FilterComposure Filter = "composure"
	// FilterAggression indicates filtering by card aggression.
	FilterAggression Filter = "aggression"
	// FilterVision indicates filtering by card vision.
	FilterVision Filter = "vision"
	// FilterAwareness indicates filtering by card awareness.
	FilterAwareness Filter = "awareness"
	// FilterCrosses indicates filtering by card crosses.
	FilterCrosses Filter = "crosses"
	// FilterPhysique indicates filtering by card physique.
	FilterPhysique Filter = "physique"
	// FilterAcceleration indicates filtering by card acceleration.
	FilterAcceleration Filter = "acceleration"
	// FilterRunningSpeed indicates filtering by card running speed.
	FilterRunningSpeed Filter = "running_speed"
	// FilterReactionSpeed indicates filtering by card reaction speed.
	FilterReactionSpeed Filter = "reaction_speed"
	// FilterAgility indicates filtering by card agility.
	FilterAgility Filter = "agility"
	// FilterStamina indicates filtering by card stamina.
	FilterStamina Filter = "stamina"
	// FilterStrength indicates filtering by card strength.
	FilterStrength Filter = "strength"
	// FilterJumping indicates filtering by card jumping.
	FilterJumping Filter = "jumping"
	// FilterBalance indicates filtering by card balance.
	FilterBalance Filter = "balance"
	// FilterTechnique indicates filtering by card technique.
	FilterTechnique Filter = "technique"
	// FilterDribbling indicates filtering by card dribbling.
	FilterDribbling Filter = "dribbling"
	// FilterBallControl indicates filtering by card ball control.
	FilterBallControl Filter = "ball_control"
	// FilterWeakFoot indicates filtering by card weak foot.
	FilterWeakFoot Filter = "weak_foot"
	// FilterSkillMoves indicates filtering by card skill moves.
	FilterSkillMoves Filter = "skill_moves"
	// FilterFinesse indicates filtering by card finesse.
	FilterFinesse Filter = "finesse"
	// FilterCurve indicates filtering by card curve.
	FilterCurve Filter = "curve"
	// FilterVolleys indicates filtering by card volleys.
	FilterVolleys Filter = "volleys"
	// FilterShortPassing indicates filtering by card short passing.
	FilterShortPassing Filter = "short_passing"
	// FilterLongPassing indicates filtering by card long passing.
	FilterLongPassing Filter = "long_passing"
	// FilterForwardPass indicates filtering by card forward pass.
	FilterForwardPass Filter = "forward_pass"
	// FilterOffense indicates filtering by card offense.
	FilterOffense Filter = "offense"
	// FilterFinishingAbility indicates filtering by card finishing ability.
	FilterFinishingAbility Filter = "finishing_ability"
	// FilterShotPower indicates filtering by card shot power.
	FilterShotPower Filter = "shot_power"
	// FilterAccuracy indicates filtering by card accuracy.
	FilterAccuracy Filter = "accuracy"
	// FilterDistance indicates filtering by card distance.
	FilterDistance Filter = "distance"
	// FilterPenalty indicates filtering by card penalty.
	FilterPenalty Filter = "penalty"
	// FilterFreeKicks indicates filtering by card free kicks.
	FilterFreeKicks Filter = "free_kicks"
	// FilterCorners indicates filtering by card corners.
	FilterCorners Filter = "corners"
	// FilterHeadingAccuracy indicates filtering by card heading accuracy.
	FilterHeadingAccuracy Filter = "heading_accuracy"
	// FilterDefence indicates filtering by card defence.
	FilterDefence Filter = "defence"
	// FilterOffsideTrap indicates filtering by card offside trap.
	FilterOffsideTrap Filter = "offside_trap"
	// FilterSliding indicates filtering by card sliding.
	FilterSliding Filter = "sliding"
	// FilterTackles indicates filtering by card tackles.
	FilterTackles Filter = "tackles"
	// FilterBallFocus indicates filtering by card ball focus.
	FilterBallFocus Filter = "ball_focus"
	// FilterInterceptions indicates filtering by card interceptions.
	FilterInterceptions Filter = "interceptions"
	// FilterVigilance indicates filtering by card vigilance.
	FilterVigilance Filter = "vigilance"
	// FilterGoalkeeping indicates filtering by card goalkeeping.
	FilterGoalkeeping Filter = "goalkeeping"
	// FilterReflexes indicates filtering by card reflexes.
	FilterReflexes Filter = "reflexes"
	// FilterDiving indicates filtering by card diving.
	FilterDiving Filter = "diving"
	// FilterHandling indicates filtering by card handling.
	FilterHandling Filter = "handling"
	// FilterSweeping indicates filtering by card sweeping.
	FilterSweeping Filter = "sweeping"
	// FilterThrowing indicates filtering by card throwing.
	FilterThrowing Filter = "throwing"
	// FilterQuality indicates filtering by card quality.
	FilterQuality Filter = "quality"
	// FilterHeight indicates filtering by card height.
	FilterHeight Filter = "height"
	// FilterWeight indicates filtering by card weight.
	FilterWeight Filter = "weight"
	// FilterDominantFoot indicates filtering by card dominant foot.
	FilterDominantFoot Filter = "dominant_foot"
	// FilterType indicates filtering by card type.
	FilterType Filter = "type"
	// FilterPrice indicates filtering by card price.
	FilterPrice Filter = "price"
	// FilterPlayerName indicates filtering by card player name.
	FilterPlayerName Filter = "player_name"
)

// SliceFilters entity for slice filters.
type SliceFilters []Filters

// Pagination defines parameters of possible cards pagination.
type Pagination string

const (
	// LimitPagination indicates the cards output limit parameter on the page.
	LimitPagination Pagination = "limit"
	// PagePagination indicates to the current output page for cards.
	PagePagination Pagination = "page"
)

// DecodingURLParameters decodes url parameters to filters entity.
func (filters *SliceFilters) DecodingURLParameters(urlQuery url.Values) error {
	for key, value := range urlQuery {
		if key == string(LimitPagination) || key == string(PagePagination) {
			continue
		}

		filter := Filters{
			Name:           "",
			Value:          value[numberPositionOfURLParameter],
			SearchOperator: "",
		}

		for k, v := range sqlsearchoperators.SearchOperators {
			if strings.HasSuffix(key, k) {
				countName := len(key) - (1 + len(k))
				filter.Name = Filter(key[:countName])
				filter.SearchOperator = v
			}
		}

		keyFilter := Filter(key)
		if keyFilter == FilterQuality || keyFilter == FilterDominantFoot || keyFilter == FilterType {
			filter.Name = Filter(key)
			filter.SearchOperator = sqlsearchoperators.EQ
		}

		if filter.Name == "" {
			return ErrInvalidFilter.New("invalid name parameter - " + key)
		}
		*filters = append(*filters, filter)
	}
	return nil
}

// Validate check of valid UTF-8 bytes and type.
func (f Filters) Validate() error {
	if f.Name == FilterTactics || f.Name == FilterPositioning || f.Name == FilterComposure || f.Name == FilterAggression ||
		f.Name == FilterVision || f.Name == FilterAwareness || f.Name == FilterCrosses || f.Name == FilterPhysique ||
		f.Name == FilterAcceleration || f.Name == FilterRunningSpeed || f.Name == FilterReactionSpeed || f.Name == FilterAgility ||
		f.Name == FilterStamina || f.Name == FilterStrength || f.Name == FilterJumping || f.Name == FilterBalance ||
		f.Name == FilterTechnique || f.Name == FilterDribbling || f.Name == FilterBallControl || f.Name == FilterWeakFoot ||
		f.Name == FilterSkillMoves || f.Name == FilterFinesse || f.Name == FilterCurve || f.Name == FilterVolleys ||
		f.Name == FilterShortPassing || f.Name == FilterLongPassing || f.Name == FilterForwardPass || f.Name == FilterOffense ||
		f.Name == FilterFinishingAbility || f.Name == FilterShotPower || f.Name == FilterAccuracy || f.Name == FilterDistance ||
		f.Name == FilterPenalty || f.Name == FilterFreeKicks || f.Name == FilterCorners || f.Name == FilterHeadingAccuracy ||
		f.Name == FilterDefence || f.Name == FilterOffsideTrap || f.Name == FilterSliding || f.Name == FilterTackles ||
		f.Name == FilterBallFocus || f.Name == FilterInterceptions || f.Name == FilterVigilance || f.Name == FilterGoalkeeping ||
		f.Name == FilterReflexes || f.Name == FilterDiving || f.Name == FilterHandling || f.Name == FilterSweeping || f.Name == FilterThrowing {
		strings.ToValidUTF8(f.Value, "")

		_, err := strconv.Atoi(f.Value)
		if err != nil {
			return ErrInvalidFilter.New("%s %s", f.Value, err)
		}
		return nil
	}

	if f.Name == FilterHeight || f.Name == FilterWeight || f.Name == FilterPrice {
		strings.ToValidUTF8(f.Value, "")

		_, err := strconv.ParseFloat(f.Value, 64)
		if err != nil {
			return ErrInvalidFilter.New("%s %s", f.Value, err)
		}
		return nil
	}

	if f.Name == FilterQuality {
		strings.ToValidUTF8(f.Value, "")

		if f.SearchOperator != sqlsearchoperators.EQ {
			return ErrInvalidFilter.New("'%s' not suitable for %s", f.SearchOperator, f.Name)
		}

		quality := Quality(f.Value)
		if quality == QualityWood || quality == QualitySilver || quality == QualityGold || quality == QualityDiamond {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, "is not an indicator of quality card")
	}

	if f.Name == FilterDominantFoot {
		strings.ToValidUTF8(f.Value, "")

		if f.SearchOperator != sqlsearchoperators.EQ {
			return ErrInvalidFilter.New("'%s' not suitable for %s", f.SearchOperator, f.Name)
		}

		dominantFoot := DominantFoot(f.Value)
		if dominantFoot == DominantFootLeft || dominantFoot == DominantFootRight {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, "is not an indicator of dominant foot card")
	}

	if f.Name == FilterType {
		strings.ToValidUTF8(f.Value, "")

		if f.SearchOperator != sqlsearchoperators.EQ {
			return ErrInvalidFilter.New("'%s' not suitable for %s", f.SearchOperator, f.Name)
		}

		filterType := Type(f.Value)
		if filterType == TypeWon || filterType == TypeBought {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, "is not an indicator of type card")
	}

	return ErrInvalidFilter.New("invalid name parameter - %s", f.Name)
}
