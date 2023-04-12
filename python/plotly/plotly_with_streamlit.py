import streamlit as st
from plotly import express as px
import pandas as pd

# Create some sample data
df = pd.DataFrame({'x': [1, 2, 3, 4, 5], 'y': [1, 4, 2, 3, 5]})

# Create a Plotly scatter plot
fig = px.scatter(df, x='x', y='y')

# Show the scatter plot using Plotly's built-in Streamlit integration
st.plotly_chart(fig)
