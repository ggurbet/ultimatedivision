// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const PlayerCard: React.FC<{ className: string; id: string; index?: number }> =
    ({ className, id, index }) =>
        <img
            className={className}
            src={`${window.location.origin}/avatars/${id}.png`}
            id={index?.toString()}
        />;

