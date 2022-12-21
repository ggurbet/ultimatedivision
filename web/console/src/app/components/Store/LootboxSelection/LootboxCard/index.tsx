// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useState } from 'react';
import { useDispatch } from 'react-redux';

import { LootboxCardQuality } from './LootboxCardQuality';

import { RegistrationPopup } from '@/app/components/common/Registration';

import { UnauthorizedError } from '@/api';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { openLootbox } from '@/app/store/actions/lootboxes';
import { LootboxStats } from '@/app/types/lootbox';
import { setCurrentUser } from '@/app/store/actions/users';
import { ToastNotifications } from '@/notifications/service';

import diamond from '@static/img/StorePage/BoxCard/diamond.svg';
import gold from '@static/img/StorePage/BoxCard/gold.svg';
import silver from '@static/img/StorePage/BoxCard/silver.svg';
import wood from '@static/img/StorePage/BoxCard/wood.svg';
import lootBox from '@static/img/StorePage/BoxContent/lootBox.svg';

import './index.scss';

export const LootboxCard: React.FC<{
    lootBoxStats: LootboxStats;
    handleOpenedLootbox: Dispatch<SetStateAction<boolean>>;
    handleLootboxSelection: Dispatch<SetStateAction<boolean>>;
    handleLootboxKeeping: Dispatch<SetStateAction<boolean>>;
}> = ({ lootBoxStats, handleOpenedLootbox, handleLootboxSelection, handleLootboxKeeping }) => {
    /** Indicates if registration required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState<boolean>(false);

    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /** Closes Registration popup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

    const dispatch = useDispatch();

    const qualities = [
        {
            name: 'Wood',
            icon: wood,
        },
        {
            name: 'Silver',
            icon: silver,
        },
        {
            name: 'Gold',
            icon: gold,
        },
        {
            name: 'Diamond',
            icon: diamond,
        },
    ];
    const boxType = lootBoxStats.type === 'Regular Box' ? 'Regular Box' : 'Cool box';

    const handleAnimation = async() => {
        try {
            await dispatch(setCurrentUser());

            handleLootboxSelection(false);
            handleLootboxKeeping(false);

            await dispatch(openLootbox({ id: lootBoxStats.id, type: lootBoxStats.type }));

            handleOpenedLootbox(true);
        } catch (error: any) {
            if (error instanceof UnauthorizedError) {
                setIsRegistrationRequired(true);

                setLocalStorageItem('IS_LOGGINED', false);

                return;
            }
            ToastNotifications.failedToOpenLootbox();
        }
    };

    if (isRegistrationRequired) {
        return <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />;
    }

    return (
        <div className={` box-card box-card${boxType === 'Regular Box' ? '--regular' : '--cool'}`}>
            <div className="box-card__wrapper">
                <div className="box-card__description">
                    <h2 className="box-card__title">{boxType}</h2>
                    <div className="box-card__quantity">
                        <span className="box-card__quantity-label">Cards</span>
                        <span className="box-card__quantity-value">{lootBoxStats.quantity}</span>
                    </div>
                    <img className="box-card__icon" src={lootBox} alt="box" />
                </div>
                <div className="box-card__qualities">
                    <h3 className="box-card__qualities__title">probability</h3>
                    {lootBoxStats.dropChance.map((item, index) =>
                        <LootboxCardQuality label={qualities[index]} chance={item} key={index} />
                    )}

                    <button className="box-card__button" onClick={handleAnimation}>
                        <span className="box-card__button-text">OPEN</span>
                        <span className="box-card__button-value">{lootBoxStats.price} coin</span>
                    </button>
                </div>
            </div>
        </div>
    );
};
