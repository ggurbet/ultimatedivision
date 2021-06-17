import rectangle
    from '../../img/MarketPlacePage/marketPlaceFilterField/rectangle.png';
import search
    from '../../img/MarketPlacePage/marketPlaceFilterField/search.png';
import star
    from '../../img/MarketPlacePage/marketPlaceFilterField/star.png';
import fut
    from '../../img/MarketPlacePage/marketPlaceFilterField/fut.png';
import eye
    from '../../img/MarketPlacePage/marketPlaceFilterField/eye.png';
import stars
    from '../../img/MarketPlacePage/marketPlaceFilterField/stars.png';
import parametres
    from '../../img/MarketPlacePage/marketPlaceFilterField/parametres.png';

const filterFieldTitles = [
    {
        title: 'Search',
        src: search
    },
    {
        title: 'Version',
        src: rectangle
    },
    {
        title: 'Positions',
        src: rectangle
    },
    {
        title: 'Nations',
        src: rectangle
    },
    {
        title: 'Leagues',
        src: rectangle
    },
    {
        title: 'WRF',
        src: rectangle
    },
    {
        title: 'Stats',
        src: rectangle
    },
    {
        title: '',
        src: star
    },
    {
        title: 'PS',
        src: fut
    },
    {
        title: 'T&S',
        src: rectangle
    },
    {
        title: '',
        src: eye
    },
    {
        title: '',
        src: stars
    },
    {
        title: 'RPP',
        src: rectangle
    },
    {
        title: '',
        src: parametres
    }
    ,
    {
        title: 'Misc',
        src: rectangle
    }
];

export const filterFieldTitlesReducer = (filterFieldTitlesState = filterFieldTitles, action) => {
    return filterFieldTitlesState;
};