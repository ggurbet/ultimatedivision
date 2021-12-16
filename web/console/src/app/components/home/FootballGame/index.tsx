// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { JoinButton } from '@components/common/JoinButton';

import footballField from '@static/img/gameLanding/main/football-field.svg';
import cards from '@static/img/gameLanding/main/cards.svg';

import './index.scss';

export const FootballGame: React.FC = () =>
    <section className="football-game">
        <picture>
            <img
                className="football-game__cards"
                src={cards}
                alt="Player cards"
            />
        </picture>
        <span className="football-game__title">ULTIMATE DIVISION</span>
        <span className="football-game__subtitle">Football P2E Game</span>
        <JoinButton />
        <picture>
            <img
                className="football-game__field"
                src={footballField}
                alt="Football field"
            />
        </picture>
    </section>;

