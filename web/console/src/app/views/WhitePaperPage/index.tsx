// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Switch } from 'react-router-dom';
import { ComponentRoutes, Route } from '@/app/router';

import './index.scss';

const WhitePaper: React.FC<{ children: ComponentRoutes[] }> = ({ children }) =>
    <section className="whitepaper">
        <div className="whitepaper__wrapper">
            <Switch>
                {children.map((route, index) =>
                    <Route
                        key={index}
                        path={route.path}
                        component={route.component}
                        exact={route.exact}
                        children={route.children}
                    />
                )
                }
            </Switch>
        </div>
    </section>;


export default WhitePaper;
