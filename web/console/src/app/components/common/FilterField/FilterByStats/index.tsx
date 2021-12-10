// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';
import { FilterFieldStatsArea, FilterFieldStatsAreaProps } from '@/app/components/common/FilterField/FilterFieldStatsArea';

import { CardsQueryParametersField } from '@/card';

// TODO: rework functionality.
export const FilterByStats: React.FC<{submitSearch: (queryParameters: CardsQueryParametersField[]) => void}> = ({ submitSearch }) => {
    /** Indicates if FilterByStats component shown. */
    const [isFilterByStatsShown, setIsFilterByStatsShown] = useState(false);

    /** Shows and closes FilterByStats component. */
    const showFilterByStats = () => {
        setIsFilterByStatsShown(isFilterByStatsShown => !isFilterByStatsShown);
    };

    /** Describes defence skills of each card. */
    const [defenceMin, setDefenceMin] = useState('');
    const [defenceMax, setDefenceMax] = useState('');

    /** Describes goalkeeping skills of each card. */
    const [goalkeepingMin, setGoalkeepingMin] = useState('');
    const [goalkeepingMax, setGoalkeepingMax] = useState('');

    /** Describes offside skills of each card. */
    const [offsideMin, setOffsideMin] = useState('');
    const [offsideMax, setOffsideMax] = useState('');

    /** Describes physique skills of each card. */
    const [physiqueMin, setPhysiqueMin] = useState('');
    const [physiqueMax, setPhysiqueMax] = useState('');

    /** Describes tactics skills of each card. */
    const [tacticsMin, setTacticsMin] = useState('');
    const [tacticsMax, setTacticsMax] = useState('');

    /** Describes technique skills of each card. */
    const [techniqueMin, setTechniqueMin] = useState('');
    const [techniqueMax, setTechniqueMax] = useState('');

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

    /** Changes min offside value. */
    const changeOffsideMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        setOffsideMin(e.target.value);
    };

    /** Changes max offside value. */
    const changeOffsideMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        setOffsideMax(e.target.value);
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
            minValue: offsideMin,
            maxValue: offsideMax,
            changeMinValue: changeOffsideMin,
            changeMaxValue: changeOffsideMax,
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
            maxValue: goalkeepingMin,
            changeMinValue: changeGoalkeepingMin,
            changeMaxValue: changeGoalkeepingMax,
        },
    ];

    /** Submits query parameters by stats. */
    const handleSubmit = async() => {
        await submitSearch([
            { 'defence_gte': defenceMin },
            { 'defence_lt': defenceMax },
            { 'goalkeeping_gte': goalkeepingMin },
            { 'goalkeeping_lt': goalkeepingMax },
            { 'offside_trap_gte': offsideMin },
            { 'offside_trap_lt': offsideMax },
            { 'physique_gte': physiqueMin },
            { 'physique_lt': physiqueMax },
            { 'tactics_gte': tacticsMin },
            { 'tactics_lt': tacticsMax },
            { 'technique_gte': techniqueMin },
            { 'technique_lt': techniqueMax },
        ]);
        showFilterByStats();
    };

    /** Clears all stats values. */
    const clearStats = () => {
        setDefenceMin('');
        setDefenceMax('');
        setGoalkeepingMin('');
        setGoalkeepingMax('');
        setOffsideMin('');
        setOffsideMax('');
        setPhysiqueMin('');
        setPhysiqueMax('');
        setTacticsMin('');
        setTacticsMax('');
        setTechniqueMin('');
        setTechniqueMax('');
    };

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStats}
            isComponentShown={isFilterByStatsShown}
            title="Stats"
        >
            <div className="filter-item__dropdown-active__stats__wrapper">
                {stats.map((stat: FilterFieldStatsAreaProps) => <FilterFieldStatsArea
                    label={stat.label}
                    minValue={stat.minValue}
                    maxValue={stat.maxValue}
                    changeMinValue={stat.changeMinValue}
                    changeMaxValue={stat.changeMaxValue}
                />)}
                <div className="filter-item__dropdown-active__stats">
                    <input
                        value="APPLY"
                        type="submit"
                        className="filter-item__dropdown-active__stats__apply"
                        onClick={handleSubmit}
                    />
                </div>
                <div className="filter-item__dropdown-active__stats">
                    <input
                        type="submit"
                        className="filter-item__dropdown-active__stats__clear"
                        value="CLEAR ALL"
                        onClick={clearStats}
                    />
                </div>
            </div>
        </FilterByParameterWrapper>
    );
};
