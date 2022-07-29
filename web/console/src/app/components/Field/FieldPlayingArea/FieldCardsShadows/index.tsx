// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { useSelector } from 'react-redux';

import { RootState } from '@/app/store';
import { Club, Formations, FormationsType, SquadCard } from '@/club';

export const FieldCardsShadows: React.FC<{ club: Club; isCardDefined: (id: string) => boolean }> = ({
    isCardDefined,
    club,
}) => {
    const formation: FormationsType = useSelector((state: RootState) => state.clubsReducer.activeClub.squad.formation);

    return (
        <div>
            {club.squadCards.map((fieldCard: SquadCard, index: number) => {
                const isDefined = isCardDefined(fieldCard.card.id);

                return (
                    <div className={`playing-area__${formation}-shadows__card`} key={index}>
                        {isDefined &&
                            <img
                                src={fieldCard.card.shadow}
                                alt="card shadow"
                                className={`playing-area__${formation}-shadows__shadow`}
                            />
                        }
                    </div>
                );
            })}
        </div>
    );
};
