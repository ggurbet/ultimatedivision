// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import { LootboxContent } from '@/app/components/Store/LootboxContent';
import { LootboxSelection } from '@/app/components/Store/LootboxSelection';
import { NftSell } from '@/app/components/Store/NftSell';

import './index.scss';

const Store = () => {
    const [isOpenedLootbox, handleOpenedLootbox] = useState(false);
    const [isLootboxSelection, handleLootboxSelection] = useState(true);
    const [isLootboxKeeping, handleLootboxKeeping] = useState(false);

    const header = document.querySelector('.ultimatedivision-navbar');

    useEffect(
        () =>
            !isLootboxSelection
                ? header?.classList.add('ultimatedivision-navbar__store-opening')
                : header?.classList.remove('ultimatedivision-navbar__store-opening'),
        [isLootboxSelection]
    );

    return (
        <section className="store">
            {isLootboxSelection ?
                <>
                    <NftSell />
                    <LootboxSelection
                        handleOpenedLootbox={handleOpenedLootbox}
                        handleLootboxSelection={handleLootboxSelection}
                        handleLootboxKeeping={handleLootboxKeeping}
                    />
                </>
                :
                <LootboxContent
                    handleOpenedLootbox={handleOpenedLootbox}
                    isOpenedLootbox={isOpenedLootbox}
                    handleLootboxSelection={handleLootboxSelection}
                    isLootboxKeeping={isLootboxKeeping}
                    handleLootboxKeeping={handleLootboxKeeping}
                />
            }
        </section>
    );
};

export default Store;
