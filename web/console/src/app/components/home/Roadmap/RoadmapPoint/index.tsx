// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { RoadmapStep } from '..';
import './index.scss';

export const RoadmapPoint: React.FC<{ item: RoadmapStep }> = ({ item }) =>
    <div className={`roadmap-point__${item.step}`}>
        <div className={'roadmap-point'}>
            <p className="roadmap-point__step">{item.step}</p>
            <ul className="roadmap-point__list">
                {item.points.map((point, _) =>
                    <li className="roadmap-point__item" key={point}>
                        <div className="roadmap-point__item__bullet"></div>
                        <span className="roadmap-point__item__description">
                            {point}
                        </span>
                    </li>
                )}
            </ul>
        </div>
    </div>;
