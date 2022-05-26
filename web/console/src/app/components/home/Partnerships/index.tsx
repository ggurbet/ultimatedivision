// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import devdao from '@static/img/gameLanding/partnerships/devdao.svg';
import boosty from '@static/img/gameLanding/partnerships/boosty.svg';
import casper from '@static/img/gameLanding/partnerships/casper.svg';
import storj from '@static/img/gameLanding/partnerships/storj.svg';
import polygon from '@static/img/gameLanding/partnerships/polygon.svg';
import chickenfish from '@static/img/gameLanding/partnerships/chickenfish.svg';
import velas from '@static/img/gameLanding/partnerships/velas.svg';

import './index.scss';

export const Partnerships: React.FC = () => {
    /** Defines logos of partner companies */
    const logos: string[] = [devdao, casper, storj, polygon, chickenfish, velas, boosty];

    return (
        <section className="partnerships">
            <div className="partnerships__wrapper">
                <h2 className="partnerships__title">Our Partnerships</h2>
                <div className="partnerships__area">
                    {logos.map((logo: string, index: number) =>
                        <div key={index} className="partnerships__area__item">
                            <div className="partnerships__area__item__wrapper">
                                <img className="partnerships__area__item__logo" key={index} src={logo} alt="logo" />
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </section>
    );
};
