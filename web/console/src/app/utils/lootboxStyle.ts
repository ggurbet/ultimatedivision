// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import coolBoxBody from '@static/img/StorePage/BoxContent/coolBoxBody.svg';
import coolBoxCover from '@static/img/StorePage/BoxContent/coolBoxCover.svg';
import regularBoxBody from '@static/img/StorePage/BoxContent/regularBoxBody.svg';
import regularBoxCover from '@static/img/StorePage/BoxContent/regularBoxCover.svg';

/** function for getting right box for animation */
export function boxStyle(length: number) {
    const REGULAR_BOX_LENGTH = 5;
    const box = {
        body: regularBoxBody,
        cover: regularBoxCover,
    };
    if (length > REGULAR_BOX_LENGTH) {
        box.body = coolBoxBody;
        box.cover = coolBoxCover;
    };

    return box;
}
