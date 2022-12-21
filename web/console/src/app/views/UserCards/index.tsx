// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useMemo, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { CardsArea } from '@components/UserCards/CardsArea';
import { FilterField } from '@components/common/FilterField';
import { FilterByPrice } from '@components/common/FilterField/FilterByPrice';
import { FilterByStats } from '@components/common/FilterField/FilterByStats';
import { FilterByStatus } from '@components/common/FilterField/FilterByStatus';
import { FilterByVersion } from '@components/common/FilterField/FilterByVersion';
import { Paginator } from '@components/common/Paginator';
import { RegistrationPopup } from '@/app/components/common/Registration';

import { UnauthorizedError } from '@/api';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { RootState } from '@/app/store';
import { createCardsQueryParameters, getCurrentCardsQueryParameters, listOfCards } from '@/app/store/actions/cards';
import { CardsQueryParametersField } from '@/card';

import './index.scss';

const UserCards: React.FC = () => {
    const { page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);
    const { currentCardsPage } = useSelector((state: RootState) => state.cardsReducer);

    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    const dispatch = useDispatch();

    const cardsQueryParameters = getCurrentCardsQueryParameters();

    /** Indicates if registration is required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    /** Exposes default page number. */
    const DEFAULT_PAGE_INDEX: number = 1;

    /** Submits search by cards query parameters. */
    const submitSearch = async(queryParameters: CardsQueryParametersField[]) => {
        createCardsQueryParameters(queryParameters);
        await dispatch(listOfCards(DEFAULT_PAGE_INDEX));
    };

    /** Closes RegistrationPopup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

    useEffect(() => {
        (async() => {
            try {
                await dispatch(listOfCards(currentCardsPage));
            } catch (error: any) {
                if (error instanceof UnauthorizedError) {
                    setIsRegistrationRequired(true);

                    setLocalStorageItem('IS_LOGGINED', false);
                }
            }
        })();
    }, []);

    return (
        <section className="user-cards">
            {isRegistrationRequired && <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />}
            <FilterField>
                <FilterByVersion submitSearch={submitSearch} cardsQueryParameters={cardsQueryParameters} />
                <FilterByStats cardsQueryParameters={cardsQueryParameters} submitSearch={submitSearch} />
                <FilterByPrice />
                <FilterByStatus />
            </FilterField>
            <CardsArea />
            <Paginator getCardsOnPage={listOfCards} itemsCount={page.totalCount} selectedPage={currentCardsPage} />
        </section>
    );
};

export default UserCards;
