// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import draws from '@static/img/profile/draws.svg';
import player from '@static/img/profile/player.svg';
import cup from '@static/img/profile/cup.svg';
import loses from '@static/img/profile/loses.svg';
import promoted from '@static/img/profile/promoted.svg';
import relegated from '@static/img/profile/relegated.svg';
import highestDivision from '@static/img/profile/highest-division.svg';

import './index.scss';

const Statistics: React.FC = () =>
    <div className="profile__statistics">
        <span className="profile__statistics__title">statistics</span>
        <div className="profile__statistics__info">
            <div className="profile__statistics__info__section">
                <img src={player} alt="Player" />
                <div className="info">
                    <span className="info__value">228</span>
                    <span className="info__title">Manager Experience
                            (seasons)
                    </span>
                </div>
            </div>
            <div className="profile__statistics__info__divider"></div>
            <div className="profile__statistics__info__section">
                <img src={cup} alt="Cup" />
                <div className="info">
                    <span className="info__value">228</span>
                    <span className="info__title">Wins</span>
                </div>
            </div>
            <div className="profile__statistics__info__section">
                <img src={loses} alt="Loses" />
                <div className="info">
                    <span className="info__value">322</span>
                    <span className="info__title">Loses</span>
                </div>
            </div>
            <div className="profile__statistics__info__section">
                <img src={draws} alt="Draws" />
                <div className="info">
                    <span className="info__value">420</span>
                    <span className="info__title">Draws</span>
                </div>
            </div>
            <div className="profile__statistics__info__divider"></div>
            <div className="profile__statistics__info__section">
                <img src={promoted} alt="Promoted logo" />
                <div className="info">
                    <span className="info__value">1</span>
                    <span className="info__title">Times promoted</span>
                </div>
            </div>
            <div className="profile__statistics__info__section">
                <img src={relegated} alt="Relegated logo" />
                <div className="info">
                    <span className="info__value">1</span>
                    <span className="info__title">Times relegated</span>
                </div>
            </div>
            <div className="profile__statistics__info__divider"></div>
            <div className="profile__statistics__info__section">
                <img src={highestDivision} alt="Highest Division" />
                <div className="info">
                    <span className="info__value">3</span>
                    <span className="info__title">Highest Division</span>
                </div>
            </div>
        </div>
    </div>;

export default Statistics;
