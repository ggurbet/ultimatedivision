// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { CardWithStats } from '@/card';

import './index.scss';

export const FootballerCardInformation: React.FC<{ card: CardWithStats }> = ({ card }) => {
    const overalInfo = card.infoBlock;

    return (
        <div className="footballer-card-information">
            {overalInfo.map((item, index) =>
                <div className="footballer-card-information__item" key={index}>
                    <div className="footballer-card-information__item-title">
                        {item.label}
                    </div>
                    <div className="footballer-card-information__item-value">
                        <>
                            {item.value}
                            {/* TODO: Need delete or rewrite this code, after backend changes */}
                            {/* <img
                                className="footballer-card-information__item-icon"
                                src={item.icon}
                                alt={item.icon && 'item icon'}
                            /> */}
                        </>
                    </div>
                </div>
            )}
        </div>
    );
};
