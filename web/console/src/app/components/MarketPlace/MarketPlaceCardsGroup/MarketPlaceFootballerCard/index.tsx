// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { Link } from 'react-router-dom';

import { Lot } from '@/marketplace';

import { PlayerCard } from '@components/common/PlayerCard';


import './index.scss';

export const MarketPlaceFootballerCard: React.FC<{ lot: Lot; place?: string }> = ({ lot }) =>
    <div
        className="marketplace-playerCard"
    >
        <Link
            className="marketplace-playerCard__link"
            to={`/lot/${lot.id}`}
        >
            <PlayerCard
                card={lot.card}
                parentClassName={'marketplace-playerCard'}
            />
        </Link>
    </div >;
