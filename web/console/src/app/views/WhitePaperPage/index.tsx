// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Route, Switch } from 'react-router-dom';
import { ComponentRoutes, RouteConfig } from '@/app/routes';

import './index.scss';

const WhitePaper: React.FC<{ children: ComponentRoutes[] }> = ({ children }) =>
    <section className="whitepaper">
        <div className="whitepaper__wrapper">
            <Switch>
                {RouteConfig.Whitepaper.children &&
                    RouteConfig.Whitepaper.children.map((route, index) =>
                        <Route
                            key={index}
                            path={route.path}
                            component={route.component}
                            exact={route.exact}
                        />
                    )
                }
            </Switch>
        </div>
    </section>;


export default WhitePaper;
