// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardStats }
    from '@components/FootballerCardPage/FootballerCardStats';

import { Card } from '@/app/types/fotballerCard';

import './index.scss';

export const FootballerCardStatsArea: React.FC<{ card: Card }> = ({ card }) => {
    const FIRST_CARD_INDEX = 0;
    const stats = card.stats;

    return (
        <div className="footballer-card-stats">
            {stats.map((item, index) =>
                <FootballerCardStats
                    key={index}
                    props={item}
                />,
            )}
        </div>
    );
};
