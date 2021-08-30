// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { BoxSelection } from '@/app/components/Store/BoxSelection';
import './index.scss';

const Store = () =>
    <section className="store">
        <h1 className="store__title">Box</h1>
        <div className="store__wrapper">
            <BoxSelection />
        </div>
    </section>;


export default Store;
