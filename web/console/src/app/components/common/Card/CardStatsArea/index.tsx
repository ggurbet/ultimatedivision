// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardStats }
    from '@/app/components/common/Card/CardStats';

import { CardWithStats } from '@/card';

import './index.scss';

export const FootballerCardStatsArea: React.FC<{ card: CardWithStats }> = ({ card }) => {
    const stats = card.statsArea;

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
