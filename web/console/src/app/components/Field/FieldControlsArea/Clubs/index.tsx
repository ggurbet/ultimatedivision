// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import Slider from 'react-slick';

import { Manager, Owner } from '@static/img/FieldPage/clubs';

import './index.scss';

const Clubs: React.FC = () => {
    const [activeClub, setActiveClub] = useState<string>('CLUB 1');
    const [activeComposition, setActiveComposition] =
        useState<string>('Composition 1');
    const [isActiveDropdown, setIsActiveDropdown] = useState<boolean>(false);
    const [clubId, setClubId] = useState<number | null>(null);

    // TODO: Mock data
    const clubs: Array<{ logo: any; name: string; whose: string }> = [
        { logo: <Owner />, name: 'CLUB 1', whose: 'owner' },
        { logo: <Owner />, name: 'CLUB 2', whose: 'owner' },
        { logo: <Owner />, name: 'CLUB 3', whose: 'owner' },
        { logo: <Manager />, name: 'CLUB 4', whose: 'manager' },
        { logo: <Manager />, name: 'CLUB 5', whose: 'manager' },
        { logo: <Manager />, name: 'CLUB 6', whose: 'manager' },
    ];

    // TODO: Mock data
    const compositions: string[] = [
        'Composition 1',
        'Composition 2',
        'Composition 3',
        'Composition 4',
    ];

    /** Method for set choosed composition to state and close dropdown block. */
    const handleChooseComposition = (composition: string) => {
        setActiveComposition(composition);
        setIsActiveDropdown(false);
    };

    /** Property for clubs slider. */
    const settings = {
        adaptiveHeight: true,
        dots: false,
        infinite: true,
        speed: 500,
        slidesToShow: 3,
        slidesToScroll: 1,
    };

    /** Show or hide helper for clubs. */
    const visabilityClubsHelper = (index: number, club: any) =>
        clubId === index &&
                <div className="club__info">
                    {club.whose === 'owner'
                        ? `are you the ${club.whose}`
                        : 'you are the manager'}
                </div>;

    return (
        <div className="field-controls-area__clubs">
            <span className="field-controls-area__clubs-name">
                {activeClub}
            </span>
            <div className="clubs">
                <Slider {...settings} className="slider">
                    {clubs.map((club, index) =>
                        <div key={index}>
                            <div
                                className={`club${
                                    club.name === activeClub ? '-active' : ''
                                }`}
                                key={index}
                                onClick={() => setActiveClub(club.name)}
                                onMouseLeave={() => setClubId(null)}
                                onMouseEnter={() => setClubId(index)}
                            >
                                {club.logo}
                                <span className="club__name">{club.name}</span>
                                {visabilityClubsHelper(index, club)}
                            </div>
                        </div>
                    )}
                </Slider>
            </div>
            <div className="composition">
                <div
                    className={`composition__choosed-item ${
                        isActiveDropdown ? 'active-dropdown' : ''
                    }`}
                    onClick={() => setIsActiveDropdown(!isActiveDropdown)}
                >
                    {activeComposition}
                </div>
                <div
                    className={`composition__list${
                        isActiveDropdown ? '-active' : ''
                    }`}
                >
                    {compositions.map((composition, index) =>
                        <div
                            className="composition__list-item"
                            key={index}
                            onClick={() => handleChooseComposition(composition)}
                        >
                            <span>{composition}</span>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default Clubs;
