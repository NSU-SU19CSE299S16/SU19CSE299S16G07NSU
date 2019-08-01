import numpy as np
import pandas as pd

dataframe = pd.read_excel(r"C:\Users\Dell\Documents\GitHub\SU19CSE299S16G07NSU\Project Code\Data\List_of_NSU_graduates\Company Details.xlsx")
print(dataframe.to_json())