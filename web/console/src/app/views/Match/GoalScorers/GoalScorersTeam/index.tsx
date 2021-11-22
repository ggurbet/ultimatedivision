// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

// TODO: Temprorary props data (we don`t know teams structure)
export const GoalScorersTeam: React.FC<any> = ({ team }) =>
    <>
        {team &&
                team.map((player: any, index: number) =>
                    <div className="player" key={index}>
                        <img
                            src={player.logo}
                            alt={`${player.name} player`}
                        ></img>
                        <span className="player__name">{player.name}</span>
                        <span className="player__goal-time">
                            {player.goalTime}
                        </span>
                        <div className="player__goals">{player.goals}</div>
                    </div>
                )}
    </>;

