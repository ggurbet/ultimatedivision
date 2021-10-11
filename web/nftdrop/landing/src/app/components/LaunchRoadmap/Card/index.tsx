// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import box from '@static/images/launchRoadmap/box1.svg';

import './index.scss';

export const Card: React.FC<{
    card: {
        title: string,
        subTitle: string,
        description: string,
        image: string,
    }
}> = ({ card }) => {

    return (
        <div className="card">
            <div className="card__text-area">
                <h1 className="card__title">
                    {card.title}
                </h1>
                <p className="card__description">
                    {card.description}
                </p>
                <div className="card__box">
                    <img
                        className="card__box__present"
                        src={box}
                        alt="utlimate box"
                    />
                    <p className="card__box__subtitle">
                        {card.subTitle}
                    </p>
                </div>
            </div>
            <img
                src={card.image}
                alt="diagram"
                className="card__image"
            />
        </div>
    );
};
