// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import PlayersList from '@/app/components/playerProfile/PlayersList';
import Statistics from '@/app/components/playerProfile/Statistics';

import profileLogo from '@static/img/profile/player-logo.png';
import owner from '@static/img/profile/owner.svg';
import manager from '@static/img/profile/manager.svg';

import './index.scss';

// TODO: Mock data (waiting for backend).
const OWNED_INDEX: number = 10;
const MANAGED_INDEX: number = 5;
// TODO: Mock data (waiting for backend).
const OwnedPlayers: string[] = new Array(OWNED_INDEX).fill('REAL MADRID');
const ManagedPlayers: string[] = new Array(MANAGED_INDEX).fill('REAL MADRID');

const PlayerProfile: React.FC = () =>
    <section className="profile">
        <div className="profile__wrapper">
            <div className="profile__info">
                <div className="profile__info__gradient"></div>
                <img className="logo" src={profileLogo} alt="Player logo" />
                <span className="player-name">player one</span>
            </div>
            <PlayersList title="Owned" logo={owner} players={OwnedPlayers} />
            <PlayersList title="Managed" logo={manager} players={ManagedPlayers} />
            <Statistics />
        </div>
    </section>;

export default PlayerProfile;
