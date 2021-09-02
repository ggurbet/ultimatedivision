// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';

import { LootboxStats } from '@/app/types/lootBox';

import { LootboxCard } from './LootboxCard';

import box from '@static/img/StorePage/BoxCard/box.svg';
import coolBox from '@static/img/StorePage/BoxCard/coolBox.svg';

import './index.scss';

export const LootboxSelection: React.FC<{ handleOpening: Dispatch<SetStateAction<boolean>> }> = ({ handleOpening }) => {
    const REGULAR_BOX_CARDS_QUANTITY = 5;
    const COOL_BOX_CARDS_QUANTITY = 10;
    /** TODO: remove test code */
    const boxesData = [
        new LootboxStats(
            '1',
            box,
            'Regular Box',
            REGULAR_BOX_CARDS_QUANTITY,
            // eslint-disable-next-line
            [80, 15, 4, 1],
            '200,000',
        ),
        new LootboxStats(
            '2',
            coolBox,
            'Cool Box',
            COOL_BOX_CARDS_QUANTITY,
            // eslint-disable-next-line
            [70, 20, 8, 2],
            '500,000',
        ),
    ];

    return (
        <div className="box-selection">
            {boxesData.map((item, index) =>
                <LootboxCard
                    data={item}
                    key={index}
                    handleOpening={handleOpening}
                />
            )}
        </div>
    );
};
