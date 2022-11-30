// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import fieldPlayerRight from '@static/img/gameLanding/gameInfo/player-right.svg';
import fieldPlayerLeft from '@static/img/gameLanding/gameInfo/player-left.svg';

import './index.scss';

export const GameInfo: React.FC = () =>
    <section className="game-info">
        <div className="game-info__content" >
            <div className="game-info__content__text game-info__content__text__left" >
                <h2 className="game-info__title">Game <span className="game-info__title__second-part">info</span></h2>
                <p className="game-info__text">Play football tactical game and compete against other
                    players on any platform.Our quick turn-based gameplay
                    allows football fans challenge each other globally.
                </p>
                <p className="game-info__text">You can form squads of your NFT players and try to outsmart
                    your opponents or you can improve your gameplay and decision-making
                    on the field. Find you our way to win.
                </p>
            </div>
            <img className="game-info__field-player" src={fieldPlayerRight} alt="field-player" />
        </div>
        <div className="game-info__content game-info__content__second-part">
            <img className="game-info__field-player" src={fieldPlayerLeft} alt="field-player" />
            <div className="game-info__content__text game-info__content__text__right">
                <p className="game-info__text">Every week, a global competition cycle updates
                    your club's standing in one of the division. Try to get promoted
                    to the Ultimate Division to get the highest weekly rewards.
                </p>
                <p className="game-info__text">Use your winnings to upgrade your
                    team and club's facilities or cash out at any time you like.
                    Look for the transfer market for lucrative deals and get to experience football
                    club economy on the blockchain.
                </p>
            </div>
        </div>

    </section>;

