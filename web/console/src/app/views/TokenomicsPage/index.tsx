// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Switch } from 'react-router-dom';
import { ComponentRoutes, Route } from '@/app/router';

import './index.scss';

const Tokenomics: React.FC<{ children: ComponentRoutes[] }> = ({ children }) =>
    <div className="tokenomics">
        <div className="tokenomics__wrapper">
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
    </div>;


export default Tokenomics;
