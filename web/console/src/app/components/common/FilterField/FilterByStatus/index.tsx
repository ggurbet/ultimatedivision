// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useDispatch } from 'react-redux';

import { FilterByParameterWrapper } from '@/app/components/common/FilterField/FilterByParameterWrapper';

import { listOfCards } from '@/app/store/actions/cards';

// TODO: rework functionality.
export const FilterByStatus: React.FC = () => {
    /** Indicates if FilterByStatus component shown. */
    const [isFilterByStatusShown, setIsFilterByStatusShown] = useState(false);

    const dispatch = useDispatch();

    /** Shows and closes FilterByStatus component. */
    const showFilterByStatus = () => {
        setIsFilterByStatusShown(isFilterByStatusShown => !isFilterByStatusShown);
    };

    /** Indicates if is choosed locked status of cards. */
    const [isLockedStatus, setIsLockedStatus] = useState<boolean>(false);
    /** Indicates if is choosed unlocked status of cards. */
    const [isUnLockedStatus, setIsUnlockedStatus] = useState<boolean>(false);

    /** Chooses locked status of cards. */
    const chooseLockedStatus = () => {
        setIsLockedStatus(isLockedStatus => !isLockedStatus);
    };

    /** Chooses unlocked status of cards. */
    const chooseUnlockedStatus = () => {
        setIsUnlockedStatus(isUnLockedStatus => !isUnLockedStatus);
    };

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** TODO: it is not added yet to query parameters on back-end. */
    /** Submits query parameters by status. */
    const handleSubmit = async() => {
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
        showFilterByStatus();
    };

    return (
        <FilterByParameterWrapper
            showComponent={showFilterByStatus}
            isComponentShown={isFilterByStatusShown}
            title="Status"
        >
            <input
                id="checkbox-locked"
                className="filter-item__dropdown-active__checkbox"
                type="checkbox"
                onClick={chooseLockedStatus}
            />
            <label
                className="filter-item__dropdown-active__text"
                htmlFor="checkbox-locked"
            >
                Locked
            </label>
            <input
                id="checkbox-unlocked"
                className="filter-item__dropdown-active__checkbox"
                type="checkbox"
                onClick={chooseUnlockedStatus}
            />
            <label
                className="filter-item__dropdown-active__text"
                htmlFor="checkbox-unlocked"
            >
                Unlocked
            </label>
            <input
                value="APPLY"
                type="submit"
                className="filter-item__dropdown-active__apply"
                onClick={handleSubmit}
            />
        </FilterByParameterWrapper>
    );
};
