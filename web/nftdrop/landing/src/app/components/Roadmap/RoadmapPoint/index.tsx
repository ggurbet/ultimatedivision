// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

export const RoadmapPoint: React.FC<{
    item: {
        date: string,
        points: string[],
        id: number,
        done: boolean
    }
}> = ({ item }) => {

    return (
        <div className="roadmap-point">
            <p className="roadmap-point__date">{item.date}</p>
            <ul className="roadmap-point__list">
                {item.points.map((point, index) => (
                    <li className="roadmap-point__item" key={index}>
                        <div className="roadmap-point__item__bullet"></div>
                        <span className="roadmap-point__item__description">
                            {point}
                        </span>
                    </li>
                ))}
            </ul>
        </div>
    );
};

