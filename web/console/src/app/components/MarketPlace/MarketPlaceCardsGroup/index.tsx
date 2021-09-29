// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Lot } from '@/marketplace';

import { MarketPlaceFootballerCard } from '@components/MarketPlace/MarketPlaceCardsGroup/MarketPlaceFootballerCard';

import './index.scss';

export const MarketPlaceCardsGroup: React.FC<{ lots: Lot[] }> = ({ lots }) =>
    <div className="marketplace-cards">
        <div className="marketplace-cards__wrapper">
            {lots.map((lot: Lot, index: number) =>
                <MarketPlaceFootballerCard lot={lot} key={index} />
            )}
        </div>
    </div>;
