celestials = GetCachedCelestials()
fleets, _ = GetFleets()
master = nil
maxVal = 0
for celestial in celestials {
    ships, _ = celestial.GetShips()
    value = ships.FleetValue()
    coord = celestial.GetCoordinate()
    for fleet in fleets {
        if (fleet.Origin.Equal(coord) && fleet.Mission != PARK) || (fleet.Destination.Equal(coord) && fleet.Mission == PARK) {
            value += fleet.Ships.FleetValue()
        }
    }
    if value > maxVal {
        maxVal = value
        master = celestial
    }
}

Print("master is", master.GetCoordinate())