// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { JoinButton } from "@components/common/JoinButton";

import cardsDesktop from "@static/img/gameLanding/main/cards-desktop.svg";
import cardsTablet from "@static/img/gameLanding/main/cards-tablet.svg";
import fieldDesktop from "@static/img/gameLanding/main/field-desktop.svg";
import fieldTablet from "@static/img/gameLanding/main/field-tablet.svg";

import "./index.scss";

export const FootballGame: React.FC = () => (
    <section className="football-game">
        <picture>
            <source media="(min-width: 1024px)" srcSet={cardsDesktop} />
            <source media="(max-width: 1023px)" srcSet={cardsTablet} />
            <img className="football-game__cards" src={cardsDesktop} alt="Player cards" />
        </picture>
        <span className="football-game__title">ULTIMATE DIVISION</span>
        <span className="football-game__subtitle">Football P2E Game</span>
        <JoinButton />
        <picture>
            <source media="(min-width: 1024px)" srcSet={fieldDesktop} />
            <source media="(max-width: 1023px)" srcSet={fieldTablet} />
            <img className="football-game__field" src={fieldDesktop} alt="Football field" />
        </picture>
    </section>
);
