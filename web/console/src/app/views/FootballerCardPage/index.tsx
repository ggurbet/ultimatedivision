// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { FootballerCardIllustrations } from '@components/FootballerCard/FootballerCardIllustrations';
import { FootballerCardPrice } from '@components/FootballerCard/FootballerCardPrice';
import { FootballerCardStatsArea } from '@components/FootballerCard/FootballerCardStatsArea';
import { FootballerCardInformation } from '@components/FootballerCard/FootballerCardInformation';

import './index.scss';
import { Card } from '@/card';

const FootballerCard: React.FC = (props) => {
    // @ts-ignore
    const cardData: Card = props.location.state.card;

    return (
        <div className="footballer-card">
            <div className="footballer-card__border">
                <div className="footballer-card__wrapper">
                    <div className="footballer-card__name-wrapper">
                        <h1 className="footballer-card__name">
                            {cardData.playerName}
                        </h1>
                    </div>
                    <FootballerCardIllustrations card={cardData} />
                    <div className="footballer-card__stats-area">
                        <FootballerCardPrice />
                        <FootballerCardStatsArea card={cardData} />
                        <FootballerCardInformation card={cardData} />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default FootballerCard;
