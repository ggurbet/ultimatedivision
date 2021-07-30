//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { FootballerCardIllustrations } from '@components/FootballerCardPage/FootballerCardIllustrations';
import { FootballerCardPrice } from '@components/FootballerCardPage/FootballerCardPrice';
import { FootballerCardStatsArea } from '@components/FootballerCardPage/FootballerCardStatsArea';
import { FootballerCardInformation } from '@components/FootballerCardPage/FootballerCardInformation';

import './index.scss';

const FootballerCard: React.FC = (props) => {
    // @ts-ignore
    const cardData = props.location.state.card;
    const FIRST_CARD_INDEX = 0;

    return (
        <div className="footballer-card">
            <div className="footballer-card__border">
                <div className="footballer-card__wrapper">
                    <div className="footballer-card__name-wrapper">
                        <h1 className="footballer-card__name">
                            {cardData.overalInfo[FIRST_CARD_INDEX].value}
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
