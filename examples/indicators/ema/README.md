# EMA (Exponential Moving Average) Example

This example demonstrates how to use the Exponential Moving Average (EMA) indicator from the gotal library.

## What it does

The EMA is a type of moving average that gives more weight to recent prices, making it more responsive to new information compared to a simple moving average. This example shows how to:

1. Create a price data series
2. Calculate the EMA from the input series
3. Output the results in multiple formats (console, CSV, and plot)

## How to run

```bash
cd examples/indicators/ema
go run main.go
```

## Output

The example will generate:

- **Console output**: Prints the EMA series values to the terminal
- **ema.csv**: A CSV file containing the EMA data series
- **ema.png**: A visual plot of the EMA data saved as a PNG image (800x600 pixels)

## Code walkthrough

```go
// Create a new price series
input := series.New("price")

// Calculate EMA from the input series
ema, err := gotal.EMA(input)

// Print results to console
ema.Print()

// Save data to CSV
ema.Save("ema.csv")

// Generate and save plot
plot := ema.Plot()
plot.Save("ema.png", 800, 600)
```
