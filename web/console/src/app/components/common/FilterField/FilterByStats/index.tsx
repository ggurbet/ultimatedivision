// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useContext } from 'react';

import { FilterFieldStatsArea, FilterFieldStatsAreaProps } from '@/app/components/common/FilterField/FilterFieldStatsArea';
import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';
import { CardsQueryParameters, CardsQueryParametersField } from '@/card';

import { FilterContext } from '../index';

// TODO: rework functionality.
export const FilterByStats: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => Promise<void>;
    cardsQueryParameters: CardsQueryParameters;
}> = ({ submitSearch, cardsQueryParameters }) => {
    const { activeFilterIndex, setActiveFilterIndex }: {
        activeFilterIndex: number;
        setActiveFilterIndex: React.Dispatch<React.SetStateAction<number>>;
    } = useContext(FilterContext);
    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const FILTER_BY_STATS_INDEX = 2;
    /** Indicates if FilterByStats component shown. */
    const [isFilterByStatsShown, setIsFilterByStatsShown] = useState(false);

    const isVisible = FILTER_BY_STATS_INDEX === activeFilterIndex && isFilterByStatsShown;

    /** Shows and closes FilterByStats component. */
    const showFilterByStats = () => {
        setActiveFilterIndex(FILTER_BY_STATS_INDEX);
        setIsFilterByStatsShown(isFilterByStatsShown => !isFilterByStatsShown);
    };

    /** Describes defence skills of each card. */
    const [defenceMin, setDefenceMin] = useState(cardsQueryParameters.defence_gte);
    const [defenceMax, setDefenceMax] = useState(cardsQueryParameters.defence_lte);

    /** Describes goalkeeping skills of each card. */
    const [goalkeepingMin, setGoalkeepingMin] = useState(cardsQueryParameters.goalkeeping_gte);
    const [goalkeepingMax, setGoalkeepingMax] = useState(cardsQueryParameters.goalkeeping_lte);

    /** Describes offense skills of each card. */
    const [offenseMin, setOffenseMin] = useState(cardsQueryParameters.offense_gte);
    const [offenseMax, setOffenseMax] = useState(cardsQueryParameters.offense_lte);

    /** Describes physique skills of each card. */
    const [physiqueMin, setPhysiqueMin] = useState(cardsQueryParameters.physique_gte);
    const [physiqueMax, setPhysiqueMax] = useState(cardsQueryParameters.goalkeeping_lte);

    /** Describes tactics skills of each card. */
    const [tacticsMin, setTacticsMin] = useState(cardsQueryParameters.tactics_gte);
    const [tacticsMax, setTacticsMax] = useState(cardsQueryParameters.tactics_lte);

    /** Describes technique skills of each card. */
    const [techniqueMin, setTechniqueMin] = useState(cardsQueryParameters.technique_gte);
    const [techniqueMax, setTechniqueMax] = useState(cardsQueryParameters.technique_lte);

    /** Changes min defence value. */
    const changeDefenceMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setDefenceMin(e.target.value);
    };

    /** Changes max defence value. */
    const changeDefenceMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setDefenceMax(e.target.value);
    };

    /** Changes min goalkeeping value. */
    const changeGoalkeepingMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setGoalkeepingMin(e.target.value);
    };

    /** Changes max goalkeeping value. */
    const changeGoalkeepingMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setGoalkeepingMax(e.target.value);
    };

    /** Changes min offense value. */
    const changeOffenseMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setOffenseMin(e.target.value);
    };

    /** Changes max offense value. */
    const changeOffenseMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setOffenseMax(e.target.value);
    };

    /** Changes min physique value. */
    const changePhysiqueMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPhysiqueMin(e.target.value);
    };

    /** Changes max physique value. */
    const changePhysiqueMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPhysiqueMax(e.target.value);
    };

    /** Changes min tactics value. */
    const changeTacticsMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTacticsMin(e.target.value);
    };

    /** Changes max tactics value. */
    const changeTacticsMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTacticsMax(e.target.value);
    };

    /** Changes min technique value. */
    const changeTechniqueMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTechniqueMin(e.target.value);
    };

    /** Changes max technique value. */
    const changeTechniqueMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTechniqueMax(e.target.value);
    };

    /** Describes stats values separated by main parameters. */
    const stats: FilterFieldStatsAreaProps[] = [
        {
            label: 'TAC',
            minValue: tacticsMin,
            maxValue: tacticsMax,
            changeMinValue: changeTacticsMin,
            changeMaxValue: changeTacticsMax,
        },
        {
            label: 'OFF',
            minValue: offenseMin,
            maxValue: offenseMax,
            changeMinValue: changeOffenseMin,
            changeMaxValue: changeOffenseMax,
        },
        {
            label: 'TEC',
            minValue: techniqueMin,
            maxValue: techniqueMax,
            changeMinValue: changeTechniqueMin,
            changeMaxValue: changeTechniqueMax,
        },
        {
            label: 'PHY',
            minValue: physiqueMin,
            maxValue: physiqueMax,
            changeMinValue: changePhysiqueMin,
            changeMaxValue: changePhysiqueMax,
        },
        {
            label: 'DEF',
            minValue: defenceMin,
            maxValue: defenceMax,
            changeMinValue: changeDefenceMin,
            changeMaxValue: changeDefenceMax,
        },
        {
            label: 'GK',
            minValue: goalkeepingMin,
            maxValue: goalkeepingMax,
            changeMinValue: changeGoalkeepingMin,
            changeMaxValue: changeGoalkeepingMax,
        },
    ];

    /** Submits query parameters by stats. */
    const handleSubmit = async() => {
        await submitSearch([
            { 'defence_gte': defenceMin },
            { 'defence_lte': defenceMax },
            { 'goalkeeping_gte': goalkeepingMin },
            { 'goalkeeping_lte': goalkeepingMax },
            { 'offense_gte': offenseMin },
            { 'offense_lte': offenseMax },
            { 'physique_gte': physiqueMin },
            { 'physique_lte': physiqueMax },
            { 'tactics_gte': tacticsMin },
            { 'tactics_lte': tacticsMax },
            { 'technique_gte': techniqueMin },
            { 'technique_lte': techniqueMax },
        ]);
        setIsFilterByStatsShown(false);
        setActiveFilterIndex(DEFAULT_FILTER_ITEM_INDEX);
    };

    /** Clears all stats values. */
    const clearStats = () => {
        setDefenceMin('');
        setDefenceMax('');
        setGoalkeepingMin('');
        setGoalkeepingMax('');
        setOffenseMin('');
        setOffenseMax('');
        setPhysiqueMin('');
        setPhysiqueMax('');
        setTacticsMin('');
        setTacticsMax('');
        setTechniqueMin('');
        setTechniqueMax('');
    };

    useEffect(() => {
        FILTER_BY_STATS_INDEX !== activeFilterIndex && setIsFilterByStatsShown(false);
    }, [activeFilterIndex]);

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStats}
            isVisible={isVisible}
            title="Stats"
        >
            <div className="filter-item__dropdown-active__stats__wrapper">
                {stats.map((stat: FilterFieldStatsAreaProps, index: number) => <FilterFieldStatsArea
                    key={index}
                    label={stat.label}
                    minValue={stat.minValue}
                    maxValue={stat.maxValue}
                    changeMinValue={stat.changeMinValue}
                    changeMaxValue={stat.changeMaxValue}
                />)}

            </div>
            <div className="filter-item__dropdown-active__stats__buttons">
                <input
                    value="APPLY"
                    type="submit"
                    className="filter-item__dropdown-active__stats__apply"
                    onClick={handleSubmit}
                />
                <input
                    type="submit"
                    className="filter-item__dropdown-active__stats__clear"
                    value="CLEAR ALL"
                    onClick={clearStats}
                />
            </div>
        </FilterByParameterWrapper>
    );
};
