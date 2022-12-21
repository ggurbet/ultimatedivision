// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useContext, useEffect, useState } from 'react';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { CardsQueryParameters, CardsQueryParametersField } from '@/card';
import { FilterContext } from '../index';

// TODO: rework functionality.
export const FilterByVersion: React.FC<{
    submitSearch: (queryParameters: CardsQueryParametersField[]) => void;
    cardsQueryParameters: CardsQueryParameters;
}> = ({ submitSearch, cardsQueryParameters }) => {
    const {
        activeFilterIndex,
        setActiveFilterIndex,
    }: {
        activeFilterIndex: number;
        setActiveFilterIndex: React.Dispatch<React.SetStateAction<number>>;
    } = useContext(FilterContext);
    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const FILTER_BY_VERSION_INDEX = 1;
    /** Indicates if FilterByVersion component shown. */
    const [isFilterByVersionShown, setIsFilterByVersionShown] = useState(false);

    const isVisible = FILTER_BY_VERSION_INDEX === activeFilterIndex && isFilterByVersionShown;

    /** Shows and closes FilterByVersion component. */
    const showFilterByVersion = () => {
        setActiveFilterIndex(FILTER_BY_VERSION_INDEX);
        setIsFilterByVersionShown((isFilterByVersionShown) => !isFilterByVersionShown);
    };

    /** Describes version parameters. */
    const [version, setVersion] = useState<string[]>(cardsQueryParameters.quality && cardsQueryParameters.quality);

    /** Indicates if is choosed diamond quality of cards. */
    const [isDiamondQuality, setIsDiamondQuality] = useState<boolean>(false);
    /** Indicates if is choosed gold quality of cards. */
    const [isGoldQuality, setIsGoldQuality] = useState<boolean>(false);
    /** Indicates if is choosed silver quality of cards. */
    const [isSilverQuality, setIsSilverQuality] = useState<boolean>(false);
    /** Indicates if is choosed wood quality of cards. */
    const [isWoodQuality, setIsWoodQuality] = useState<boolean>(false);

    /** Chooses diamond quality of cards. */
    const chooseDiamondQuality = () => {
        setIsDiamondQuality((isDiamondQuality) => !isDiamondQuality);
    };

    /** Chooses gold quality of cards. */
    const chooseGoldQuality = () => {
        setIsGoldQuality((isGoldQuality) => !isGoldQuality);
    };

    /** Chooses silver quality of cards. */
    const chooseSilverQuality = () => {
        setIsSilverQuality((isSilverQuality) => !isSilverQuality);
    };

    /** Chooses wood quality of cards. */
    const chooseWoodQuality = () => {
        setIsWoodQuality((isWoodQuality) => !isWoodQuality);
    };

    /** Changes quality of cards. */
    const changeQuality: () => string[] = () => {
        const qualities: string[] = [];

        if (isDiamondQuality) {
            qualities.push('diamond');
        }

        if (isGoldQuality) {
            qualities.push('gold');
        }

        if (isSilverQuality) {
            qualities.push('silver');
        }

        if (isWoodQuality) {
            qualities.push('wood');
        }

        return qualities;
    };

    /** Submits query parameters by quality. */
    const handleSubmit = async() => {
        await submitSearch([{ quality: changeQuality() }]);
        setIsFilterByVersionShown(false);
        setActiveFilterIndex(DEFAULT_FILTER_ITEM_INDEX);
    };

    /** Checks current versions. */
    const checkCurrentVersion = () => {
        // TODO: rework functionality.
        setIsDiamondQuality(Boolean(version && version.includes('diamond')));
        setIsGoldQuality(Boolean(version && version.includes('gold')));
        setIsSilverQuality(Boolean(version && version.includes('silver')));
        setIsWoodQuality(Boolean(version && version.includes('wood')));
    };

    useEffect(() => {
        FILTER_BY_VERSION_INDEX !== activeFilterIndex && setIsFilterByVersionShown(false);
        setVersion(cardsQueryParameters.quality);
        checkCurrentVersion();
    }, [activeFilterIndex, cardsQueryParameters]);

    return (
        <FilterByParameterWrapper showComponent={showFilterByVersion} isVisible={isVisible} title="Version">
            <div className="filter-item__dropdown-active__wrapper">
                <div className="filter-item__dropdown-active__switcher">
                    <input
                        id="division-checkbox-wood"
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        checked={isWoodQuality}
                        onChange={chooseWoodQuality}
                    />
                    <label className="filter-item__dropdown-active__slider" htmlFor="division-checkbox-wood"></label>
                    <p className="filter-item__dropdown-active__text">wood</p>
                </div>
                <div className="filter-item__dropdown-active__switcher">
                    <input
                        id="checkbox-silver"
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        checked={isSilverQuality}
                        onChange={chooseSilverQuality}
                    />

                    <label className="filter-item__dropdown-active__slider" htmlFor="checkbox-silver"></label>
                    <p className="filter-item__dropdown-active__text">silver</p>
                </div>
                <div className="filter-item__dropdown-active__switcher">
                    <input
                        id="checkbox-gold"
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        checked={isGoldQuality}
                        onChange={chooseGoldQuality}
                    />

                    <label className="filter-item__dropdown-active__slider" htmlFor="checkbox-gold"></label>
                    <p className="filter-item__dropdown-active__text">gold</p>
                </div>
                <div className="filter-item__dropdown-active__switcher">
                    <input
                        id="checkbox-diamond"
                        className="filter-item__dropdown-active__checkbox"
                        type="checkbox"
                        checked={isDiamondQuality}
                        onChange={chooseDiamondQuality}
                    />
                    <label className="filter-item__dropdown-active__slider" htmlFor="checkbox-diamond"></label>
                    <p className="filter-item__dropdown-active__text">diamond</p>
                </div>
            </div>
            <input value="APPLY" type="submit" className="filter-item__dropdown-active__apply" onClick={handleSubmit} />
        </FilterByParameterWrapper>
    );
};
