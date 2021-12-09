// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { ClubCardsArea } from '@components/Club/ClubCardsArea';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';
import { RegistrationPopup } from '@/app/components/common/Registration/Registration';

import { UnauthorizedError } from '@/api';
import { listOfCards } from '@/app/store/actions/cards';
import { RootState } from '@/app/store';

import './index.scss';

const Club: React.FC = () => {
    const dispatch = useDispatch();
    const { page } = useSelector((state: RootState) => state.cardsReducer.cardsPage);

    /** Describes default page number. */
    const DEFAULT_PAGE_NUMBER: number = 1;

    /** Indicates if registration is required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    /** Closes RegistrationPopup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

    useEffect(() => {
        (async() => {
            try {
                await dispatch(listOfCards(DEFAULT_PAGE_NUMBER));
            } catch (error: any) {
                if (error instanceof UnauthorizedError) {
                    setIsRegistrationRequired(true);

                    return;
                };
            };
        })();
    }, []);

    return (
        <section className="club">
            {isRegistrationRequired && <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />}
            <h1 className="club__title">
                MY CARDS
            </h1>
            <FilterField />
            <ClubCardsArea />
            <Paginator
                getCardsOnPage={listOfCards}
                itemsCount={page.totalCount}
                selectedPage={page.currentPage}
            />
        </section>
    );
};

export default Club;
