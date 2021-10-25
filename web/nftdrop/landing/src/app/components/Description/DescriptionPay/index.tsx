// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { AnimationImage } from '@components/common/AnimationImage';

import playToEarnData from '@static/images/description/playToEarn/data.json';

import animationImage_0 from '@static/images/description/playToEarn/images/img_0.png';
import animationImage_1 from '@static/images/description/playToEarn/images/img_1.png';
import animationImage_2 from '@static/images/description/playToEarn/images/img_2.png';
import animationImage_3 from '@static/images/description/playToEarn/images/img_3.png';
import animationImage_4 from '@static/images/description/playToEarn/images/img_4.png';

import './index.scss';

export const DescriptionPay = () => {
    const animationImages: string[] = [
        animationImage_0,
        animationImage_1,
        animationImage_2,
        animationImage_3,
        animationImage_4,
    ];

    return (
        <div className="description-pay">
            <AnimationImage
                className={'description-pay__radar'}
                heightFrom={1000}
                heightTo={-500}
                loop={false}
                animationData={playToEarnData}
                animationImages={animationImages}
                isNeedScrollListener={true}
            />
            <div className="description-pay__text-area">
                <h2 className="description-pay__title">Play-to-Earn</h2>
                <p className="description-pay__text">
                    Club Owners who hold a Founder Collection NFT will be
                    awarded the in-game title of UD Founder. The UD Founders
                    will receive exclusive airdrops and will start the game in
                    UD’s top division.
                </p>
            </div>
        </div>
    );
};
