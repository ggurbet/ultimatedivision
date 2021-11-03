// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import realMadrid from '@static/img/divisions/realmadrid.png';

import { DivisionClub } from '@/app/types/division';

import './index.scss';

const Division: React.FC = () => {
    /** TODO: replase test datas. Just for test rigth now. */
    const divisionClub: DivisionClub = {
        position: '1',
        club: {
            name: 'Real Madrid',
            icon: realMadrid,
        },
        games: '30',
        wins: '20',
        draws: '6',
        defeats: '5',
        goalDifference: '+50',
        points: '50',
    };

    /** TODO: delete this. Rigth now for test */
    const CLUBS_COUNT: number = 50;
    const divisionClubs = new Array(CLUBS_COUNT).fill(divisionClub);

    /** Variables describes upper and lower clubs position
     * in general division table. */
    const COEFFICIENT: number = 0.1;

    const UPPER_BREAKPOINT: number
        = Math.round(divisionClubs.length * COEFFICIENT);
    const LOWER_BREAKPOINT: number
        = divisionClubs.length - UPPER_BREAKPOINT;

    const titles: string[] = [
        '#',
        'Club',
        'Games',
        'Wins',
        'Draws',
        'Defeats',
        'The difference of goals conceded',
        'Points',
    ];

    const changeGradationDivisionClassName = (position: string) => {
        let className: string = '';

        if (+position <= UPPER_BREAKPOINT) {
            className = '-upper';
        };

        if (+position >= LOWER_BREAKPOINT) {
            className = '-lower';
        };

        return className;
    };

    return (
        <section className="division">
            <div className="division__titles">
                <h1 className="division__titles__main">
                    Division 1
                </h1>
                <span className="division__titles__count">
                    {CLUBS_COUNT}
                    <span className="division__titles__count__text">
                        Teams
                    </span>
                </span>
            </div>
            <div className="division__filters">
                <span className="division__filters__title">
                    Filters
                </span>
                <div className="division__filters__item">
                    <input
                        id="division-checkbox-1"
                        className="division__filters__item__checkbox"
                        type="checkbox"
                    />
                    <label
                        className="division__filters__item__text"
                        htmlFor="division-checkbox-1"
                    >
                        played all the matches
                    </label>
                </div>
                <div className="division__filters__item">
                    <input
                        id="division-checkbox-2"
                        className="division__filters__item__checkbox"
                        type="checkbox"
                    />
                    <label
                        className="division__filters__item__text"
                        htmlFor="division-checkbox-2"
                    >
                        did not play in the current draw
                    </label>
                </div>
            </div>
            <table className="division__clubs">
                <thead>
                    <tr className="division__clubs__titles">
                        {titles.map((title: string, index: number) =>
                            <th
                                key={index}
                                className="division__clubs__titles__item"
                            >
                                {title}
                            </th>
                        )}
                    </tr>
                </thead>
                <tbody>
                    {divisionClubs.map((divisionClub: DivisionClub, index: number) =>
                        <tr
                            className={`division__clubs__club${changeGradationDivisionClassName(divisionClub.position)}`}
                            key={index}
                        >
                            <td className="division__clubs__club__item">
                                <span className="division__clubs__club__item__position">
                                    {divisionClub.position}
                                </span>
                            </td>
                            <td className="division__clubs__club__item">
                                <img
                                    src={divisionClub.club.icon}
                                    className="division__clubs__club__item__icon"
                                />
                                {divisionClub.club.name}
                            </td>
                            <td className="division__clubs__club__item">
                                {divisionClub.games}
                            </td>
                            <td className="division__clubs__club__item">
                                {divisionClub.wins}
                            </td>
                            <td className="division__clubs__club__item">
                                {divisionClub.draws}
                            </td>
                            <td className="division__clubs__club__item">
                                {divisionClub.defeats}
                            </td>
                            <td className="division__clubs__club__item">
                                {divisionClub.goalDifference}
                            </td>
                            <td className="division__clubs__club__item">
                                {divisionClub.points}
                            </td>
                        </tr>
                    )}
                </tbody>
            </table>
        </section>
    );
};

export default Division;
