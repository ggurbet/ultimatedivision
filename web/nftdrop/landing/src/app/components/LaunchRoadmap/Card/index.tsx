// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { AnimationImage } from '@components/common/AnimationImage';

import box from '@static/images/launchRoadmap/box.svg';

import './index.scss';

export const Card: React.FC<{
    card: {
        id: string;
        title: string;
        subTitle: string;
        description: string;
        animation: any;
    };
}> = ({ card }) => {
    return (
        <div className="card">
            <div className="card__text-area">
                <h1 className="card__title">{card.title}</h1>
                <p className="card__description">{card.description}</p>
                <div className="card__box">
                    <img
                        className="card__box__present"
                        src={box}
                        alt="utlimate box"
                    />
                    <p className="card__box__subtitle">{card.subTitle}</p>
                </div>
            </div>
            <AnimationImage
                className={`card__image-${card.id}`}
                heightFrom={1000}
                heightTo={-400}
                loop={true}
                animationData={card.animation}
                animationImages={[]}
                isNeedScrollListener={true}
            />
        </div>
    );
};
