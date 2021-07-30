//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import './index.scss';

const GameMechanics = () =>
        <div className="game-mechanics">
                <h1 className="game-mechanics__title">
                        Game Mechanics
                </h1>
                <h2 className="game-mechanics__subtitle">
                        Match Day
                </h2>
                <p className="game-mechanics__description">
                        Players may initiate a PvP game by creating a smart-contract.In order to play, each user (club owner or dedicated manager) needs to pick his starting 11 players and up to 6 substitute players.
                        <br /><br />
                        Once a match is initiated, the user can see his opponent’s squad and can proceed to determining the playing formation, positions and detailed instructions of each player.
                        <br /><br />
                        After both users have confirmed their tactics (or when the time expires) the game result is calculated based on the strength of the squads and tactics utilised. While the result of the game is determined by the smart-contract and probabilities, the visual interface provides game simulation.
                </p>
                <h2 className="game-mechanics__subtitle">
                        Ranking and Divisions
                </h2>
                <p className="game-mechanics__description">
                        Each Football club is assigned to a division ranked from 1 (highest) to 10 (lowest). Upon creation, each club starts from division 10. Clubs can only be matched up with clubs from the same division.
                        <br /><br />
                        At the end of every week, 10% of best-performing clubs get promoted to a higher division, and 10% of the worst-performing clubs get relegated to a lower division.If a club is inactive during the week, it’s division can not be changed. However, every club has to play at least 3 matches during the week to get weekly rewards.
                </p>
                <h2 className="game-mechanics__subtitle">
                        Managers and club owners
                </h2>
                <p className="game-mechanics__description">
                        Owning a football club requires investment (buying the NFTs). However, a new player can start his path by working as the manager.
                        <br />
                        A manager can be hired via a smart-contract to operate the club for a fixed period of time in exchange for a proportion of the club’s proceeds.
                </p>
        </div>;


export default GameMechanics;
