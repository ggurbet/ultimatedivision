// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';

import './index.scss';

export const CardMintingProgress: React.FC<{
    max: number;
    activeCardsCount: number;
}> = ({ max, activeCardsCount }) => {
    /** Creates array with empty elements and then fill them */
    const [cards, setCards] = useState(new Array(max).fill({}));

    const fillActiveCards = () => {
        setCards(
            cards.map((_, index) => ({
                active: index < activeCardsCount,
            }))
        );
    };

    useEffect(() => {
        fillActiveCards();
    }, [activeCardsCount]);

    return (
        <div className="card-minting">
            {cards.map((card, index) =>
                <div key={index} className={`card-minting__item ${card.active ? 'card-minting__active-item' : ''}`}></div>
            )}
        </div>
    );
};
