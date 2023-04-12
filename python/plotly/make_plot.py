import plotly.graph_objs as go
import pandas as pd

# Create some sample data
df = pd.DataFrame({'x': [1, 2, 3, 4, 5], 'y': [1, 4, 2, 3, 5]})

# Create a Plotly figure with a line plot
fig = go.Figure(data=go.Scatter(x=df['x'], y=df['y'], mode='lines'))

# Customize the figure layout
fig.update_layout(
    title='Line Plot Example',
    xaxis_title='X Axis Label',
    yaxis_title='Y Axis Label'
)

# Display the plot
fig.show()
