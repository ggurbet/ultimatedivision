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

/** Domain entity Partner implementation */
class Partner {
    /** default partner implementation */
    constructor(public name: string = '', public logo: string = '') {}
}

export const Partnerships: React.FC = () => {
    /** Defines partners companies */
    const partners: Partner[] = [
        new Partner('polygon', polygon),
        new Partner('velas', velas),
        new Partner('casper', casper),
        new Partner('devdao', devdao),
        new Partner('storj', storj),
        new Partner('boosty', boosty),
        new Partner('chickenfish', chickenfish),
    ];

    return (
        <section className="partnerships">
            <div className="partnerships__wrapper">
                <h2 className="partnerships__title">Our <span className="partnerships__title__second-part">Partners</span></h2>
                <div className="partnerships__area">
                    {partners.map((partner: Partner, _) =>
                        <div key={partner.logo} className="partnerships__area__item">
                            <div className="partnerships__area__item__wrapper">
                                <img className={`partnerships__area__item__logo partnerships__area__item__logo__${partner.name}`} src={partner.logo} alt="logo" />
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </section>
    );
};
