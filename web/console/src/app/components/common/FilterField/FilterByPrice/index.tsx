// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useContext, useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { listOfCards } from '@/app/store/actions/cards';
import { FilterContext } from '../index';

// TODO: rework functionality.
export const FilterByPrice: React.FC = () => {
    const { activeFilterIndex, setActiveFilterIndex }: {
        activeFilterIndex: number;
        setActiveFilterIndex: React.Dispatch<React.SetStateAction<number>>;
    } = useContext(FilterContext);

    /** Exposes default index which does not exist in array. */
    const DEFAULT_FILTER_ITEM_INDEX = -1;
    const FILTER_BY_PRICE_INDEX = 3;
    /** Indicates if FilterByPrice component shown. */
    const [isFilterByPriceShown, setIsFilterByPriceShown] = useState(false);

    const isVisible = FILTER_BY_PRICE_INDEX === activeFilterIndex && isFilterByPriceShown;

    const dispatch = useDispatch();

    /** Shows and closes FilterByPrice component. */
    const showFilterByPrice = () => {
        setActiveFilterIndex(FILTER_BY_PRICE_INDEX);
        setIsFilterByPriceShown(isFilterByPriceShown => !isFilterByPriceShown);
    };

    /** Indicates min price value. */
    const [minPrice, setMinPrice] = useState<string>('');
    /** Indicates max price value. */
    const [maxPrice, setMaxPrice] = useState<string>('');

    /** Changes min price value for cards. */
    const changeMinPrice = (e: React.ChangeEvent<HTMLInputElement>) => {
        setMinPrice(e.target.value);
    };

    /** Changes max price value for cards. */
    const changeMaxPrice = (e: React.ChangeEvent<HTMLInputElement>) => {
        setMaxPrice(e.target.value);
    };

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** TODO: it is not added yet to query parameters on back-end. */
    /** Submits query parameters by status. */
    const handleSubmit = async() => {
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
        setIsFilterByPriceShown(false);
        setActiveFilterIndex(DEFAULT_FILTER_ITEM_INDEX);
    };

    useEffect(() => {
        FILTER_BY_PRICE_INDEX !== activeFilterIndex && setIsFilterByPriceShown(false);
    }, [activeFilterIndex]);

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByPrice}
            isVisible={isVisible}
            title="Price"
        >
            <div className="filter-item__dropdown-active__wrapper">
                <input
                    placeholder="Min"
                    className="filter-item__dropdown-active__price"
                    type="text"
                    onChange={changeMinPrice}
                />
                <input
                    placeholder="Max"
                    className="filter-item__dropdown-active__price"
                    type="text"
                    onChange={changeMaxPrice}
                />
            </div>
            <input
                value="APPLY"
                type="submit"
                className="filter-item__dropdown-active__apply"
                onClick={handleSubmit}
            />
        </FilterByParameterWrapper>
    );
};
