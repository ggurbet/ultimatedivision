// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';

import './index.scss';

// The maximum length of the list of players that is shown on the page.
const MAX_PLAYERS_LENGTH: number = 5;
// TODO: Mock data (so that the list of players does not start from 0)
const INCREMENT_INDEX: number = 1;

const PlayersList: React.FC<{title: string; logo: string; players: string[]}> = ({ title, logo, players }) =>
{
    const [playersLength, setPlayersLength] = useState<null | number>(players.length);
    const [isFullPlayersBlockVisible, setIsFullPlayersBlockVisible] = useState<boolean>(false);

    /** Changing the button class depending on the length of the player list. */
    const showMoreBtnClassName: string = playersLength && playersLength > MAX_PLAYERS_LENGTH ? '' : '-disable';
    /** Changing the players block class depending on the button "Show more" click. */
    const playersBlockClassName: string = isFullPlayersBlockVisible ? 'open' : 'hidden';

    return (
        <div className="profile__player-list">
            <div className="profile__player-list__title">
                <img className="logo" src={logo} alt={`${logo} logo`} />
                <span className="title">Clubs {title} by Player</span>
            </div>
            <div className={`players-wrapper-${playersBlockClassName}`}>
                {players && players.map((player: string, index: number) =>
                    <div className="profile__player-list__item" key={index}>
                        <div className="profile__player-list__item__number">
                            <span className="profile__player-list__item__number-text">{index + INCREMENT_INDEX}</span>
                        </div>
                        <span className="profile__player-list__item__name">{player}</span>
                    </div>
                )}
            </div>
            <button className={`profile__player-list__btn${showMoreBtnClassName} ${playersBlockClassName}`}
                onClick={() => !showMoreBtnClassName && setIsFullPlayersBlockVisible(!isFullPlayersBlockVisible)}>
                <span className="btn-text">show more</span>
            </button>
        </div>
    );
};

export default PlayersList;
