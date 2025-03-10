// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { lazy } from 'react';
import { Route, Switch, useLocation } from 'react-router-dom';

const MarketPlace = lazy(() => import('@/app/views/MarketPlacePage'));
const UserCards = lazy(() => import('@/app/views/UserCards'));
const Card = lazy(() => import('@/app/views/CardAndLotPage/CardPage'));
const Lot = lazy(() => import('@/app/views/CardAndLotPage/LotPage'));
const Field = lazy(() => import('@/app/views/FieldPage'));
const WhitePaper = lazy(() => import('@/app/views/WhitePaperPage'));
const Tokenomics = lazy(() => import('@/app/views/TokenomicsPage'));
const Store = lazy(() => import('@/app/views/StorePage'));
const Division = lazy(() => import('@/app/views/Division'));
const Match = lazy(() => import('@/app/views/Match'));
const MatchFinder = lazy(() => import('@components/common/MatchFinder'));
const Home = lazy(() => import('@/app/views/Home'));
const Navbar = lazy(() => import('@/app/components/common/Navbar'));
const FootballGame = lazy(() => import('@/app/views/FootballGame'));
const PlayerProfile = lazy(() => import('@/app/views/PlayerProfile'));
const AuthWrapper = lazy(() => import('@/app/components/common/Registration/AuthWrapper'));

import Summary from '@components/WhitePaper/Summary';
import GameMechanics from '@components/WhitePaper/GameMechanics';
import PayToEarnEconomy from '@components/WhitePaper/PayToEarnEconomy';
import Technology from '@components/WhitePaper/Technology';
import Fund from '@components/Tokenomics/Fund';
import PlayToEarn from '@components/Tokenomics/PlayToEarn';
import Spending from '@components/Tokenomics/Spending';
import Staking from '@components/Tokenomics/Staking';

/** Route base config implementation */
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: any,
        public exact: boolean,
        public className?: string,
        public children?: ComponentRoutes[]
    ) { }
    /** Method for creating child subroutes path */
    public with(child: ComponentRoutes, parrent: ComponentRoutes): ComponentRoutes {
        child.path = `${parrent.path}/${child.path}`;

        return this;
    }
    /** Call with method for each child */
    public addChildren(children: ComponentRoutes[]): ComponentRoutes {
        this.children = children.map((item) => item.with(item, this));

        return this;
    }
}

/** Route config implementation */
export class RouteConfig {
    public static FootballGame: ComponentRoutes = new ComponentRoutes('/game', FootballGame, true);
    public static MarketPlace: ComponentRoutes = new ComponentRoutes('/marketplace', MarketPlace, true, 'page__marketplace');
    public static Lot: ComponentRoutes = new ComponentRoutes('/lot/:id', Lot, true);
    public static Card: ComponentRoutes = new ComponentRoutes('/card/:id', Card, false);
    public static Division: ComponentRoutes = new ComponentRoutes(
        /** TODO: it will be replaced with id parameter */
        '/divisions',
        Division,
        true
    );
    public static Home: ComponentRoutes = new ComponentRoutes('/home', Home, true, 'page__home');
    public static Field: ComponentRoutes = new ComponentRoutes('/field', Field, true);
    public static Store: ComponentRoutes = new ComponentRoutes('/store', Store, true, 'page__store');
    public static Cards: ComponentRoutes = new ComponentRoutes('/cards', UserCards, true, 'page__cards');
    public static Match: ComponentRoutes = new ComponentRoutes('/match', Match, true);
    public static PlayerProfile: ComponentRoutes = new ComponentRoutes('/player-profile', PlayerProfile, true);
    public static Whitepaper: ComponentRoutes = new ComponentRoutes('/whitepaper', WhitePaper, false, 'page__whitepaper');
    public static Tokenomics: ComponentRoutes = new ComponentRoutes('/tokenomics', Tokenomics, false, 'page__whitepaper');
    public static Summary: ComponentRoutes = new ComponentRoutes('summary', Summary, true, 'page__whitepaper');
    public static GameMechanics: ComponentRoutes = new ComponentRoutes('game-mechanics', GameMechanics, true, 'page__whitepaper');
    public static PayToEarnEconomy: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn-and-economy',
        PayToEarnEconomy,
        true, 'page__whitepaper'
    );
    public static Technology: ComponentRoutes = new ComponentRoutes('technology', Technology, true, 'page__whitepaper');
    public static Spending: ComponentRoutes = new ComponentRoutes('udt-spending', Spending, true, 'page__whitepaper');
    public static PayToEarn: ComponentRoutes = new ComponentRoutes('pay-to-earn', PlayToEarn, true, 'page__whitepaper');
    public static Staking: ComponentRoutes = new ComponentRoutes('staking', Staking, true, 'page__whitepaper');
    public static Fund: ComponentRoutes = new ComponentRoutes('ud-dao-fund', Fund, true, 'page__whitepaper');

    public static routes: ComponentRoutes[] = [
        RouteConfig.Home,
        RouteConfig.Field,
        RouteConfig.MarketPlace,
        RouteConfig.Cards,
        RouteConfig.Card,
        RouteConfig.Division,
        RouteConfig.Lot,
        RouteConfig.Store,
        RouteConfig.Match,
        RouteConfig.FootballGame,
        RouteConfig.PlayerProfile,
        RouteConfig.Whitepaper.addChildren([
            RouteConfig.Summary,
            RouteConfig.GameMechanics,
            RouteConfig.PayToEarnEconomy,
            RouteConfig.Technology,
        ]),
        RouteConfig.Tokenomics.addChildren([
            RouteConfig.Spending,
            RouteConfig.PayToEarn,
            RouteConfig.Staking,
            RouteConfig.Fund,
        ]),
    ];
}

/** Route config that implements auth actions */
export class AuthRouteConfig {
    public static AuthWrapper: ComponentRoutes = new ComponentRoutes('/auth-velas', AuthWrapper, true);
    public static Default: ComponentRoutes = new ComponentRoutes('/', Home, true, 'page__home');

    public static routes: ComponentRoutes[] = [AuthRouteConfig.Default, AuthRouteConfig.AuthWrapper];
}

export const Routes = () => {
    const FILTERED_IS_PAGE_CLASSNAME = 0;

    const location = useLocation();
    const currentLocation = location.pathname;

    const pageClassName = RouteConfig.routes.filter((route, _) =>
        route.className && route.path === currentLocation ? route.className : ''
    )[FILTERED_IS_PAGE_CLASSNAME]?.className || AuthRouteConfig.routes.filter((route, _) =>
        route.className && route.path === currentLocation ? route.className : ''
    )[FILTERED_IS_PAGE_CLASSNAME]?.className;

    return (
        <div className={`page ${pageClassName ? pageClassName : ''}`}>
            <Switch>
                {AuthRouteConfig.routes.map((route, index) =>
                    <Route key={index} path={route.path} component={route.component} exact={route.exact} />
                )}
                <Route>
                    <Navbar />
                    <MatchFinder />
                    {RouteConfig.routes.map((route, index) =>
                        <Route key={index} path={route.path} component={route.component} exact={route.exact} />
                    )}
                </Route>
            </Switch>
        </div>
    );
};
