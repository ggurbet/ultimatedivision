// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Switch } from 'react-router-dom';
import { ComponentRoutes, Route } from '@/app/routes';

import { AboutMenu } from '../../AboutMenu';

import './index.scss';

export const Tokenomics: React.FC<{ children: ComponentRoutes[] }> = ({ children }) => {
    return (
        <div className="tokenomics">
            <AboutMenu />
            <div className="tokenomics__wrapper">
                <Switch>
                    {children.map((route, index) => (
                        <Route
                            key={index}
                            path={route.path}
                            component={route.component}
                            exact={route.exact}
                            children={route.children}
                        />
                    ))
                    }
                </Switch>
            </div>
        </div>
    )
}
