// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketPlaceFootballerCard } from '@components/MarketPlace/MarketPlaceCardsGroup/MarketPlaceFootballerCard';

import { Lot } from '@/marketplace';

import './index.scss';

export const MarketPlaceCardsGroup: React.FC<{ lots: Lot[]; handleShowModal: (lot: Lot) => void }> = ({ lots, handleShowModal }) =>
    <div className="marketplace-cards">
        <div className="marketplace-cards__wrapper">
            {lots.map((lot: Lot, index: number) =>
                <MarketPlaceFootballerCard lot={lot} key={index} handleShowModal={handleShowModal} />
            )}
        </div>
    </div>;


