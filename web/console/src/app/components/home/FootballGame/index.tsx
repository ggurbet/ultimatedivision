// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { JoinButton } from '@components/common/JoinButton';

import field from '@static/img/gameLanding/main/field.webp';
import fieldAnimation from '@static/img/gameLanding/main/field-animation.webp';

import './index.scss';

export const FootballGame: React.FC = () =>
    <section className="football-game">
        <div className="football-game__info">
            <h1 className="football-game__title">ULTIMATE <span className="football-game__title__second-part">DIVISION</span></h1>
            <h2 className="football-game__subtitle">Football P2E Game</h2>
            <div className="football-game__button-desktop"><JoinButton /></div>
        </div>
        <picture className="football-game__field">
            <img className="football-game__field__bg" src={field} alt="Football field" />
            <img className="football-game__field__animation" src={fieldAnimation} alt="Football field" />
        </picture>
        <div className="football-game__button-mobile"><JoinButton /></div>
    </section>;

