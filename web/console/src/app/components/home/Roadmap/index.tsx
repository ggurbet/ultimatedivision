// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { RoadmapPoint } from './RoadmapPoint';

import roadmap from '@static/img/gameLanding/roadmap/roadmap.png';

import './index.scss';

/** Domain entity RoadmapStep implementation */
export class RoadmapStep {
    /** default RoadmapStep implementation */
    constructor(
        public id: string = '',
        public step: string = '',
        public points: string[] = []) { }
}

export const Roadmap: React.FC = () => {
    const roadmapSteps = [
        new RoadmapStep('1', 'start', [
            'Ultimate Division Metaverse MVP launch',
            'Play to Earn mechanics available to players',
            'UDT (Ultimate Division Token) introduced',
        ]),
        new RoadmapStep('2', 'middle', [
            'Game in full swing',
            'Management roles available',
            'Club owners can tokenize their clubs and sell shares to other players',
            'UDT partnership with Top-5 Leagues',
        ]),
        new RoadmapStep('3', 'finish', [
            'Advanced gameplay introduced',
            'Local competitions launched',
            'DAO governance adopted',
        ]),
    ];

    return (
        <section className="roadmap">
            <h2 className="roadmap__title">Our <span className="roadmap__title__second-part">Roadmap</span></h2>
            {roadmapSteps.map((item: RoadmapStep) =>
                <RoadmapPoint key={item.id} item={item} />
            )}
            <img className="roadmap__image" src={roadmap} alt="roadmap" />
        </section>
    );
};
