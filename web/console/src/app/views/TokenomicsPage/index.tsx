// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Route, Switch } from 'react-router-dom';
import { ComponentRoutes, RouteConfig } from '@/app/routes';

import './index.scss';

const Tokenomics: React.FC<{ children: ComponentRoutes[] }> = ({ children }) =>
    <section className="tokenomics">
        <div className="tokenomics__wrapper">
            <Switch>
                {RouteConfig.Tokenomics.children &&
                    RouteConfig.Tokenomics.children.map((route, index) =>
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


export default Tokenomics;
