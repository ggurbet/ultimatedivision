// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import fieldPlayerRight from '@static/img/gameLanding/gameInfo/player-right.svg';
import fieldPlayerLeft from '@static/img/gameLanding/gameInfo/player-left.svg';

import './index.scss';

const GameInfo: React.FC = () =>
    <div className="game-info">
        <div className="game-info__content" >
            <div className="game-info__content__text game-info__content__text__left" >
                <h2 className="game-info__title">Game <span className="game-info__title__second-part">info</span></h2>
                <p className="game-info__text">Through simple, secure, and scalable technology, NEAR empowers
                        millions to invent and explore new experiences. Business, creativity,
                        and community are being reimagined for a more sustainable and inclusive
                        future.
                </p>
                <p className="game-info__text">Through simple, secure, and scalable technology, NEAR empowers millions to
                        invent and explore new experiences. Business, creativity, and community
                        are being reimagined for a more sustainable and inclusive future.
                </p>
            </div>
            <img className="game-info__field-player" src={fieldPlayerRight} alt="field-player"/>
        </div>
        <div className="game-info__content game-info__content__second-part">
            <img className="game-info__field-player" src={fieldPlayerLeft} alt="field-player" />
            <div className="game-info__content__text game-info__content__text__right">
                <p className="game-info__text">Through simple, secure, and scalable technology, NEAR empowers
                        millions to invent and explore new experiences.
                        Business, creativity, and community are being reimagined
                        for a more sustainable and inclusive future.
                </p>
                <p className="game-info__text">Through simple, secure, and scalable technology,
                        NEAR empowers millions to invent and explore new experiences.
                        Business, creativity, and community are being reimagined for
                        a more sustainable and inclusive future.
                </p>
                <p className="game-info__text">Through simple, secure, and scalable technology, NEAR empowers millions to
                        invent and explore new experiences. Business, creativity, and community
                        are being reimagined for a more sustainable and inclusive future.
                </p>
            </div>
        </div>

    </div>;

export default GameInfo;
