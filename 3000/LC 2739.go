func distanceTraveled(mainTank int, additionalTank int) int {
    return (mainTank + min(additionalTank, (mainTank - 1) / 4)) * 10
}
