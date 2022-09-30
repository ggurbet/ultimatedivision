// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

export const LootboxCardQuality: React.FC<{
    label: { name: string; icon: string };
    chance: number;
}> = ({ label, chance }) =>
    <div className="box-card-quality">
        <img className="box-card-quality__icon" src={label.icon} alt="quality icon" />
        <span className="box-card-quality__text">{label.name}</span>
        <span className="box-card-quality__value">{`-${chance}%`}</span>
    </div>;

