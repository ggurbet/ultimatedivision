// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { MarketplaceLot } from '@/card';
import { MarketPlaceFootballerCard } from '@components/MarketPlace/MarketPlaceCardsGroup/MarketPlaceFootballerCard';

import './index.scss';

export const MarketPlaceCardsGroup: React.FC<{ lots: MarketplaceLot[] }> = ({ lots }) =>
    <div className="marketplace-cards">
        <div className="marketplace-cards__wrapper">
            {lots.map((lot, index) =>
                <MarketPlaceFootballerCard card={lot.card} key={index} />
            )}
        </div>
    </div>;
